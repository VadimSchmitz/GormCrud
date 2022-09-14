import { Field } from "formik";

export default function FormikField({ id, label, placeholder, type }) {
  return (
    <div className="flex-col mb-3 w-auto">
      <label className="block text-gray-700 text-sm font-bold mb-[1px]" htmlFor={id}>
        {label}
      </label>
      <Field
        className="shadow appearance-none border rounded py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        id={id}
        name={id}
        placeholder={placeholder}
        type={type}
      />
    </div>
  );
}
