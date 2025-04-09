package export

import (
	"context"
	"errors"
	"fmt"

	provisioning "github.com/grafana/grafana/pkg/apis/provisioning/v0alpha1"
	"github.com/grafana/grafana/pkg/registry/apis/provisioning/jobs"
	"github.com/grafana/grafana/pkg/registry/apis/provisioning/repository"
	"github.com/grafana/grafana/pkg/registry/apis/provisioning/resources"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/dynamic"
)

func ExportFolders(ctx context.Context, repoName string, options provisioning.ExportJobOptions, folderClient dynamic.ResourceInterface, repositoryResources resources.RepositoryResources, progress jobs.JobProgressRecorder) error {
	// Load and write all folders
	// FIXME: we load the entire tree in memory
	progress.SetMessage(ctx, "read folder tree from API server")

	tree := resources.NewEmptyFolderTree()
	if err := resources.ForEach(ctx, folderClient, func(item *unstructured.Unstructured) error {
		if tree.Count() >= resources.MaxNumberOfFolders {
			return errors.New("too many folders")
		}

		// FIXME: repoName should be part of skip folder export
		return tree.AddUnstructured(item, repoName)
	}); err != nil {
		return fmt.Errorf("load folder tree: %w", err)
	}

	progress.SetMessage(ctx, "write folders to repository")
	err := repositoryResources.EnsureFolderTreeExists(ctx, options.Branch, options.Path, tree, func(folder resources.Folder, created bool, err error) error {
		result := jobs.JobResourceResult{
			Action:   repository.FileActionCreated,
			Name:     folder.ID,
			Resource: resources.FolderResource.Resource,
			Group:    resources.FolderResource.Group,
			Path:     folder.Path,
			Error:    err,
		}

		if !created {
			result.Action = repository.FileActionIgnored
		}

		progress.Record(ctx, result)
		if err := progress.TooManyErrors(); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("write folders to repository: %w", err)
	}

	return nil
}
