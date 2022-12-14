import "tw-elements";

export default function MovieEntry({ id, title, rating, year, summary }) {
  return (
    <div className="mb-1">
      <div>
        <div className="bg-white">
          <div
            className="accordion-button collapsed file:relative flex items-center w-full py-6 px-5 text-base text-gray-800 text-left bg-white transition focus:outline-none"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target={`#flush-${id}`}
            aria-expanded="false"
            aria-controls={`flush-${id}`}>
            <div className="flex-1">
              <button className="flex link pr-2 font-bold text-left">
                {title}
                <p className="my-auto ml-1 font-normal">({year})</p>
              </button>
              <div className="flex">
                <svg
                  aria-hidden="true"
                  className="mt-[1px] w-5 h-5 text-yellow-400"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                  xmlns="http://www.w3.org/2000/svg">
                  <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path>
                </svg>
                <p className="my-auto w-full">{rating} /10</p>
              </div>
            </div>
          </div>
          <div id={`flush-${id}`} className="accordion-collapse collapse" aria-labelledby="flush-1">
            <div className="py-4 px-5">
              <h2 className="font-bold text-[#424242]">Plot:</h2>
              {summary}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
