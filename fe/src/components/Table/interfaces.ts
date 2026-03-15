import type { SortOrder } from '../../core/enums.ts';
import type { PlainRecord, TableColumnDefinition } from '../../core/interfaces.ts';

export interface Props {
  columns: TableColumnDefinition[];
  rows: PlainRecord[];
  sortBy: string;
  sortOrder: SortOrder;
  loading?: boolean;
  emptyLabel: string;
}

export interface RowProps {
  item: PlainRecord;
  columns: TableColumnDefinition[];
}
