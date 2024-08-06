import { Carousel } from "./ui/Carousel";
import { Consultation } from "./ui/Consultation";
import { FAQ } from "./ui/FAQ";
import { Features } from "./ui/Features";
import { Hero } from "./ui/Hero";
import { Services } from "./ui/Services";

export const HomePage = () => {
  return (
    <>
      <Hero />
      <Carousel />
      <FAQ />
      <Features />
      <Services />
      <Consultation />
    </>
  );
};
