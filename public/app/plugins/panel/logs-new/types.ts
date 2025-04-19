import { Grammar } from 'prismjs';

import { CoreApp, DataFrame } from '@grafana/data';
import { LogListControlOptions } from 'app/features/logs/components/panel/LogList';

type onNewLogsReceivedType = (allLogs: DataFrame[], newLogs: DataFrame[]) => void;
type onLogOptionsChangeType = (option: keyof LogListControlOptions, value: string | boolean | string[]) => void;

export function isOnNewLogsReceivedType(callback: unknown): callback is onNewLogsReceivedType {
  return typeof callback === 'function';
}

export function isOnLogOptionsChange(callback: unknown): callback is onLogOptionsChangeType {
  return typeof callback === 'function';
}

export function isLogsGrammar(grammar: unknown): grammar is Grammar {
  return grammar !== null && typeof grammar === 'object' && Object.getPrototypeOf(grammar) === Object.prototype;
}

export function isCoreApp(app: unknown): app is CoreApp {
  return (
    app === CoreApp.CloudAlerting ||
    app === CoreApp.Correlations ||
    app === CoreApp.Dashboard ||
    app === CoreApp.Explore ||
    app === CoreApp.PanelEditor ||
    app === CoreApp.PanelViewer ||
    app === CoreApp.UnifiedAlerting ||
    app === CoreApp.Unknown
  );
}
