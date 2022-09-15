import { useState } from "react";
import Pagination from "./Pagination";
import PaginatedMovies from "./MovieEntry/PaginatedMovies";

export default function Content({ movies }) {

  const [currentPage, setCurrentPage] = useState(1);
  const [recordsPerPage] = useState(5);

  const indexOfLastRecord = currentPage * recordsPerPage;
  const indexOfFirstRecord = indexOfLastRecord - recordsPerPage;
  const currentRecords = movies.slice(indexOfFirstRecord, indexOfLastRecord);
  const nPages = Math.ceil(movies.length / recordsPerPage);

  return (
    <div>
      <h1> Movies </h1>
      <PaginatedMovies data={currentRecords} />
      <Pagination
        className="flex"
        nPages={nPages}
        currentPage={currentPage}
        setCurrentPage={setCurrentPage}
      />
    </div>
  );
}
