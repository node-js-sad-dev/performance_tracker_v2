import { Lap } from '../../core/interfaces.ts';
import { Entity } from '../../core/enums.ts';

export interface Props {
  entity: Entity;
}

export interface State {
  laps: Array<Lap>;
}
