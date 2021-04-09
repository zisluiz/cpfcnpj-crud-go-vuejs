import Query from "@/model/Query";

/*
Convert query object into query url params
*/
const toQuery = (params: any, query: Query) => {
  query.page = params.first;
  query.resultsPerPage = params.rows;

  if (params.sortField != null) {
    query.sorts = [];
    query.sorts.push({
      name: params.sortField,
      order: params.sortOrder == 1 ? "asc" : "desc",
    });
  } else {
    query.sorts = null;
  }
};

export default toQuery;
