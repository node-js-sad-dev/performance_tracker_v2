import { CURRENT_VIEW } from '../../core/consts.ts';

export interface Props {
  currentView: keyof typeof CURRENT_VIEW;
}
