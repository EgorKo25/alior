import { Header } from "src/widgets/Header";
import { Footer } from "src/widgets/Footer";
import { Outlet } from "react-router-dom";

export const Layout = () => {
  return (
    <>
      <Header />

      <Outlet />

      <Footer />
    </>
  );
};
