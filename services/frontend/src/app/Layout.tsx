import { Header } from "src/widgets/Header";
import { Footer } from "src/widgets/Footer";

export const Layout = ({ children }: React.PropsWithChildren) => {
  return (
    <>
      <Header />

      {children}

      <Footer />
    </>
  );
};
