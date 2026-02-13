import { Entity } from './enums.ts';

export const EntityConfig = {
  [Entity.CARS]: {
    route: '/cars',
    headerTitle: 'Cars Telemetry',
  },
  [Entity.LAPS]: {
    route: '/laps',
    headerTitle: 'Laps Telemetry',
  },
  [Entity.TRACKS]: {
    route: '/tracks',
    headerTitle: 'Tracks Telemetry',
  },
  [Entity.GAMES]: {
    route: '/games',
    headerTitle: 'Games Telemetry',
  },
  [Entity.GEAR]: {
    route: '/gear',
    headerTitle: 'Gear Telemetry',
  },
};
