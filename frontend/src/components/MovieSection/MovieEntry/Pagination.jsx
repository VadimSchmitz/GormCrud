const Pagination = ({ nPages, currentPage, setCurrentPage }) => {
  const pageNumbers = [...Array(nPages + 1).keys()].slice(1);

  const nextPage = () => {
    sessionStorage.setItem("scrollPosition", window.pageYOffset);
    if (currentPage !== nPages) setCurrentPage(currentPage + 1);
  };
  const prevPage = () => {
    sessionStorage.setItem("scrollPosition", window.pageYOffset);
    if (currentPage !== 1) setCurrentPage(currentPage - 1);
  };

  return (
    <div>
      <div className="flex justify-center">
        <nav>
          <ul className="flex list-style-none">
            <li className="page-item">
              <button
                onClick={prevPage}
                className="page-link relative block py-1.5 px-3 border-0 bg-transparent outline-none transition-all duration-300 rounded text-gray-800 hover:text-gray-800 hover:bg-gray-200 focus:shadow-none">
                Prev
              </button>
            </li>
            {console.log({ pageNumbers })}
            {pageNumbers.map((pageNumber) => (
              <li
                onClick={() => setCurrentPage(pageNumber)}
                className={`${
                  currentPage == pageNumber
                    ? "page-link relative block py-1.5 px-3 border-0 outline-none transition-all duration-300 rounded bg-blue-600 text-white hover:text-white hover:bg-blue-600 shadow-md focus:shadow-md"
                    : "page-link relative block py-1.5 px-3 border-0 outline-none transition-all duration-300 rounded text-gray-800 hover:text-gray-800 hover:bg-gray-200 focus:shadow-none "
                }`}
                key={pageNumber}>
                <button>
                  {pageNumber} <span className="visually-hidden">(current)</span>
                </button>
              </li>
            ))}

            <li className="page-item">
              <button
                onClick={nextPage}
                className="page-link relative block py-1.5 px-3 border-0 bg-transparent outline-none transition-all duration-300 rounded text-gray-800 hover:text-gray-800 hover:bg-gray-200 focus:shadow-none">
                Next
              </button>
            </li>
          </ul>
        </nav>
      </div>
    </div>
  );
};

export default Pagination;
