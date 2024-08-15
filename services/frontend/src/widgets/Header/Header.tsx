import { useState, useEffect } from "react";
import { Navbar } from "./ui/Navbar";
import styles from "./Header.module.scss";
import clsx from "clsx";
import { IconBars } from "src/shared/ui/IconBars";
import { IconCross } from "src/shared/ui/IconCross";
import { MainButton } from "src/shared/ui/MainButton";

export const Header = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  // const [loading, setLoading] = useState(true);
  // const [isAtTop, setIsAtTop] = useState(true);

  // useEffect(() => {
  //   if (window.scrollY !== 0) {
  //     setIsAtTop(false);
  //   }

  //   const handleScroll = () => {
  //     if (!loading) {
  //       setIsAtTop(window.scrollY === 0);
  //     }
  //   };

  //   window.addEventListener("scroll", handleScroll);

  //   return () => {
  //     window.removeEventListener("scroll", handleScroll);
  //   };
  // }, [loading]);

  // useEffect(() => {
  //   if (isAtTop) {
  //     const timer = setTimeout(() => {
  //       setLoading(false);
  //     }, 3000);

  //     return () => clearTimeout(timer);
  //   }
  // }, [isAtTop]);

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
      {/* {loading && isAtTop && (
        <div
          className={clsx(" preloader", {
            " preloader_logo_visible": loading && isAtTop,
          })}
        >
          <img
            className={clsx(" ", {
              " animated_logo_visible": loading && isAtTop,
            })}
            src="/images/logo.svg"
            alt="Alior"
          />
        </div>
      )} */}
      <div>
        <img
          // className={clsx("", {
          //   [styles.logo]: loading && isAtTop,
          // })}
          src="./images/logo.svg"
          alt="Alior"
        />
      </div>
      <div className="hidden lg:flex">
        <Navbar isMobile={false} />
      </div>
      <div className="hidden sm:flex sm:ml-auto lg:ml-0">
        <MainButton
          className={` btn-accent ${isMenuOpen ? " hidden" : ""}`}
          title="На консультацию"
          type="submit"
          colorSchema=" btn-accent-white"
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
            colorSchema=" btn-accent-white"
          />
        </div>
      )}
    </header>
  );
};
