import type { Lap } from '../../core/interfaces.ts';

export interface Props {
  currentView: string;
}

export interface State {
  laps: Array<Lap>;
}
