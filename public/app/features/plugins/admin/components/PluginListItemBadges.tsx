import { PluginSignatureBadge, Stack } from '@grafana/ui';

import { isPluginUpdatable } from '../helpers';
import { CatalogPlugin } from '../types';

import {
  PluginEnterpriseBadge,
  PluginDisabledBadge,
  PluginInstalledBadge,
  PluginUpdateAvailableBadge,
  PluginDeprecatedBadge,
} from './Badges';

type PluginBadgeType = {
  plugin: CatalogPlugin;
};

export function PluginListItemBadges({ plugin }: PluginBadgeType) {
  // Currently renderer plugins are not supported by the catalog due to complications related to installation / update / uninstall.
  const canUpdate = isPluginUpdatable(plugin);
  if (plugin.isEnterprise) {
    return (
      <Stack height="auto" wrap="wrap">
        <PluginEnterpriseBadge plugin={plugin} />
        {plugin.isDisabled && <PluginDisabledBadge error={plugin.error} />}
        {canUpdate && <PluginUpdateAvailableBadge plugin={plugin} />}
      </Stack>
    );
  }

  return (
    <Stack height="auto" wrap="wrap">
      <PluginSignatureBadge status={plugin.signature} />
      {plugin.isDisabled && <PluginDisabledBadge error={plugin.error} />}
      {plugin.isDeprecated && <PluginDeprecatedBadge />}
      {plugin.isInstalled && <PluginInstalledBadge />}
      {canUpdate && <PluginUpdateAvailableBadge plugin={plugin} />}
    </Stack>
  );
}
