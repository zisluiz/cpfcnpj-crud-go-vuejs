export default interface Query {
  page: number;
  resultsPerPage: number;
  sorts: SortField[] | null;
  filters: FilterField[] | null;
  toParam(): any;
}

export interface SortField {
  name: string;
  order: string;
}

export interface FilterField {
  name: string;
  value: any;
}

export const emptyQuery: Query = {
  page: 0,
  resultsPerPage: 10,
  sorts: null,
  filters: null,
  toParam() {
    return {
      page: this.page,
      resultsPerPage: this.resultsPerPage,
      sorts: this.sorts != null ? JSON.stringify(this.sorts) : null,
      filters: this.filters != null ? JSON.stringify(this.filters) : null,
    };
  },
};
