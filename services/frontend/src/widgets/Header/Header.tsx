import gsap from "gsap";
import { ScrollTrigger } from "gsap/all";
import { useState, useEffect } from "react";
import { HashLink } from "react-router-hash-link";
import clsx from "clsx";
import { IconBars } from "src/shared/ui/IconBars";
import { IconCross } from "src/shared/ui/IconCross";
import { MainButton } from "src/shared/ui/MainButton";
import { Navbar } from "./ui/Navbar";
import { LogoSVG } from "../../shared/ui/LogoSVG";
import styles from "./Header.module.scss";

export const Header = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const handleResize = () => {
    if (window.innerWidth > 640 && isMenuOpen) {
      setIsMenuOpen(false);
    }
  };

  useEffect(() => {
    window.addEventListener("resize", handleResize);

    return () => {
      window.removeEventListener("resize", handleResize);
    };
  }, [isMenuOpen]);

  useEffect(() => {
    gsap.registerPlugin(ScrollTrigger);

    gsap.to(".header-logo", {
      scale: 1,
      left: "5%",
      top: "28px",
      delay: 1,
      duration: 1.25,
      transform: "translate(0,0)",
    });

    gsap.to(".white-bg-under-logo", {
      opacity: 0,
      backgroundColor: "transparent",
      delay: 1,
      duration: 1.5,
    });
  }, []);

  return (
    <header
      className={clsx(styles.header) + ` sm:justify-normal  lg:justify-between`}
    >
      <div className="white-bg-under-logo fixed top-0 left-0 w-screen h-screen bg-white pointer-events-none select-none z-40"></div>
      <div className="w-[111px]">
        <LogoSVG
          className={
            "header-logo absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 scale-[3] md:scale-[4.5] lg:scale-[6] z-50"
          }
        />
      </div>
      <div className="hidden lg:flex">
        <Navbar isMobile={false} />
      </div>
      <div className="hidden sm:flex sm:ml-auto lg:ml-0">
        <HashLink
          smooth
          to={"#consult"}
          className={`${isMenuOpen ? " hidden" : ""}`}
        >
          <MainButton
            className={` btn-accent `}
            title="На консультацию"
            type="submit"
            colorSchema=" btn-accent-white"
          />
        </HashLink>
      </div>
      <div
        className={clsx(styles.menuIcon) + ` sm:ml-16 lg:hidden `}
        onClick={() => setIsMenuOpen(!isMenuOpen)}
      >
        {isMenuOpen ? (
          // SVG для закрытого состояния меню
          <IconCross />
        ) : (
          // SVG для открытого состояния меню
          <IconBars />
        )}
      </div>
      {isMenuOpen && (
        <div className={styles.menu}>
          <Navbar isMobile={true} />
          <HashLink smooth to={"#consult"}>
            <MainButton
              className={` btn-accent`}
              title="Консультация"
              type="submit"
              colorSchema=" btn-accent-white"
            />
          </HashLink>
        </div>
      )}
    </header>
  );
};
