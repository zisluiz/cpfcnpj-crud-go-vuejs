export default interface PaginatedResponse {
  page: number;
  resultsPerPage: number;
  totalResults: number;
  content: any;
  validations: any;
}

export const emptyResponse: PaginatedResponse = {
  page: 0,
  resultsPerPage: 0,
  totalResults: 0,
  content: [],
  validations: null,
};
