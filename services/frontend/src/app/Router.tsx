import { createBrowserRouter } from "react-router-dom";
import { HomePage } from "src/pages/HomePage";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <HomePage />,
  },
]);
