import { Formik, Form } from "formik";
import FormikField from "./FormikField";
import * as Yup from "yup";

export default function MovieForm({ setMovies }) {
  const movieSchema = Yup.object().shape({
    imdb_id: Yup.string()
      .required("IMDb ID is required")
      .typeError("you must specify a text")
      .min(9, "The IMDb ID must be exactly 9 digits")
      .max(9, "The IMDb ID must be exactly 9 digits")
      .matches(/(tt[0-9]*)/g, "Must be a valid IMDb ID")
      .required("IMDb ID is required")
  });
  return (
    <div>
      <h1>Add a movie</h1>
      <p>Just give the IMDb id abd we will handle the rest.</p>
      <Formik
        initialValues={{
          imdb_id: ""
        }}
        validationSchema={movieSchema}
        onSubmit={async (values) => {
          await new Promise((r) => setTimeout(r, 20));
          const requestOptions = {
            method: "POST",
            mode: "cors",
            body: JSON.stringify(values)
          };
          fetch("http://localhost:8090/movieid", requestOptions)
            .then((response) => response.json())
            .then((data) => {
              setMovies((movies) => [data, ...movies]);
            });
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
