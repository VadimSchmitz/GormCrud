import { useState, useEffect } from "react";
import Pagination from "./MovieEntry/Pagination";
import PaginatedMovies from "./MovieEntry/PaginatedMovies";

export default function Content({ movies }) {
  const [currentPage, setCurrentPage] = useState(1);
  const [recordsPerPage] = useState(100);

  const indexOfLastRecord = currentPage * recordsPerPage;
  const indexOfFirstRecord = indexOfLastRecord - recordsPerPage;
  const currentRecords = movies.slice(indexOfFirstRecord, indexOfLastRecord);
  const nPages = Math.ceil(movies.length / recordsPerPage);

  useEffect(() => {
    const scrollPosition = sessionStorage.getItem("scrollPosition");
    if (scrollPosition) {
      window.scrollTo(0, parseInt(scrollPosition, 10));
      sessionStorage.removeItem("scrollPosition");
    }
  }, [currentPage]);

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
