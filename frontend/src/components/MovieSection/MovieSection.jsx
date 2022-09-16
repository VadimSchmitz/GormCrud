import { useState } from "react";
import Pagination from "./MovieEntry/Pagination";
import PaginatedMovies from "./MovieEntry/PaginatedMovies";

export default function Content({ movies }) {
  const [currentPage, setCurrentPage] = useState(1);
  const [recordsPerPage] = useState(250);

  const indexOfLastRecord = currentPage * recordsPerPage;
  const indexOfFirstRecord = indexOfLastRecord - recordsPerPage;
  const currentRecords = movies.slice(indexOfFirstRecord, indexOfLastRecord);
  const nPages = Math.ceil(movies.length / recordsPerPage);

  return (
    <div>
      <h1> Movies </h1>
      <PaginatedMovies itemsPerPage={4} data={currentRecords} />
      <Pagination
        className="flex"
        nPages={nPages}
        currentPage={currentPage}
        setCurrentPage={setCurrentPage}
        pageRangeDisplayed={3}
        marginPagesDisplayed={2}
        breakLabel="..."
        breakClassName="page-item"
      />
    </div>
  );
}
