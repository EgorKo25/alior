import { useState, useEffect } from "react";
import { HashLink } from "react-router-hash-link";
import clsx from "clsx";
import { IconBars } from "src/shared/ui/IconBars";
import { IconCross } from "src/shared/ui/IconCross";
import { MainButton } from "src/shared/ui/MainButton";
import { Navbar } from "./ui/Navbar";
import { LogoSVG } from "./ui/LogoSVG";
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

  return (
    <header
      className={clsx(styles.header) + ` sm:justify-normal  lg:justify-between`}
    >
      <div>
        <LogoSVG />
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
