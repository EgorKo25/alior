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
    <header className={`` + clsx(styles.header)}>
      <div className={clsx(styles.heading_logo)}>
        <img src="/images/logo.svg" alt="Alior" />
      </div>
      <div className="hidden sm:flex">
        <Navbar isMobile={false} />
      </div>
      <div className="hidden sm:flex">
        <MainButton
          className={clsx(styles.button_black_type)}
          title="На консультацию"
          type="submit"
          textColor="#fff"
        />
      </div>
      <div
        className={clsx(styles.menuIcon, { [styles.hidden]: !isMenuOpen })}
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
            className={clsx(styles.button_black_type)}
            title="Консультация"
            type="submit"
            textColor="#fff"
          />
        </div>
      )}
    </header>
  );
};
