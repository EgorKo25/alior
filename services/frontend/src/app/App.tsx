import { RouterProvider } from "react-router-dom";
import { router } from "./Router";
import "./App.css";

export const App = () => {
  return (
    <>
      <RouterProvider router={router} />
    </>
  );
};
