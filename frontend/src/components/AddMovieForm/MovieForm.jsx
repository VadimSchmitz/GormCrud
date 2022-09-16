import { Formik, Form } from "formik";
import FormikField from "./FormikField";
import * as Yup from "yup";

export default function MovieForm({ setMovies }) {
  const currentYear = new Date().getFullYear();

  const movieSchema = Yup.object().shape({
    imdb_id: Yup.string()
      .required("IMDb ID is required")
      .typeError("you must specify a text")
      .min(9, "The IMDb ID must be exactly 9 digits")
      .max(9, "The IMDb ID must be exactly 9 digits"),
    title: Yup.string().required("Title is required").typeError("you must specify a text"),
    rating: Yup.number()
      .required("Rating is required")
      .typeError("you must specify a number")
      .min(0, "Rating must be between 0 and 10")
      .max(10, "Rating must be between 0 and 10"),
    year: Yup.number()
      .required("Year is required")
      .typeError("you must specify a number")
      .min(1888, `Year must be between 1888 and ${currentYear + 5}`)
      .max(currentYear + 5, `Year must be between 1888 and ${currentYear + 5}`),
    summary: Yup.string().required("Summary is required").typeError("you must specify a text")
  });

  return (
    <div>
      <h1>Add a movie</h1>
      <Formik
        initialValues={{
          imdb_id: "",
          title: "",
          rating: "",
          year: "",
          summary: ""
        }}
        validationSchema={movieSchema}
        onSubmit={async (values) => {
          await new Promise((r) => setTimeout(r, 20));
          const requestOptions = {
            method: "POST",
            mode: "cors",
            body: JSON.stringify(values)
          };
          fetch("http://localhost:8090/movies", requestOptions)
            .then((response) => response.json())
            .then(setMovies((movies) => [values, ...movies]))
            .then(console.log("Movie added"));
        }}>
        {({ errors, touched }) => (
          <Form className="flex-col">
            <FormikField
              id="imdb_id"
              label="IMDb_id"
              placeholder="tt0993846"
              type=""
              error={errors.imdb_id}
              touched={touched.imdb_id}
            />
            <FormikField
              id="title"
              label="Title"
              placeholder="The Wolf of Wall Street"
              type=""
              error={errors.title}
              touched={touched.title}
            />

            <FormikField
              id="rating"
              label="Rating"
              placeholder="8.2"
              type=""
              error={errors.rating}
              touched={touched.rating}
            />

            <FormikField
              id="year"
              label="Year"
              placeholder="2013"
              type=""
              error={errors.year}
              touched={touched.year}
            />

            <FormikField
              id="summary"
              label="Summary"
              placeholder="Based on the true story of Jordan Belfort, from his rise to a wealthy stock-broker living the high life to his fall involving crime, corruption and the federal government."
              type=""
              error={errors.summary}
              touched={touched.summary}
            />

            <button
              className="max-w-[160px] py-2 px-3 bg-blue-500 hover:bg-blue-700 text-white font-bold"
              type="submit">
              Add a movie
            </button>
          </Form>
        )}
      </Formik>
    </div>
  );
}
