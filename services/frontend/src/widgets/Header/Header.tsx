import { useState, useEffect } from "react";
import { Navbar } from "./ui/Navbar";
import styles from "./Header.module.scss";
import clsx from "clsx";
import { IconBars } from "src/shared/ui/IconBars";
import { IconCross } from "src/shared/ui/IconCross";
import { MainButton } from "src/shared/ui/MainButton";

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

  return (
    <header
      className={clsx(styles.header) + ` sm:justify-normal  lg:justify-between`}
    >
      <div className={clsx(styles.heading_logo) + ``}>
        <img src="/images/logo.svg" alt="Alior" />
      </div>
      <div className="hidden lg:flex">
        <Navbar isMobile={false} />
      </div>
      <div className="hidden sm:flex sm:ml-auto lg:ml-0">
        <MainButton
          className={` btn-accent ${isMenuOpen ? " hidden" : ""}`}
          title="На консультацию"
          type="submit"
          circleClassName=" circle-hover-accent"
        />
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
          <MainButton
            className={` btn-accent`}
            title="Консультация"
            type="submit"
            circleClassName=" circle-hover-accent"
          />
        </div>
      )}
    </header>
  );
};
