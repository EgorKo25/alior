import { useState, useEffect } from "react";
import { NavBarItem } from "./NavBarItem";
import styles from "./Navbar.module.scss";
import mobileStyles from "./MobileNavbar.module.scss";

const menuItems = ["О нас", "Услуги", "Кейсы"];
const backgroundImages = [
  `url(chair.svg)`,
  `url(services.svg)`,
  `url(play.svg)`,
];

export const Navbar = ({ isMobile }: { isMobile?: boolean }) => {
  const [activeIndex, setActiveIndex] = useState(0);

  useEffect(() => {
    const savedIndex = localStorage.getItem("activeIndex");
    if (savedIndex !== null) {
      setActiveIndex(JSON.parse(savedIndex));
    }
  }, []);

  useEffect(() => {
    localStorage.setItem("activeIndex", JSON.stringify(activeIndex));
  }, [activeIndex]);

  const handleItemClick = (index: number) => {
    setActiveIndex(index);
  };

  const containerStyle = isMobile
    ? mobileStyles.container
    : styles.container + " bg-accent";
  const listStyle = isMobile ? mobileStyles.list : styles.list;
  return (
    <div className={containerStyle}>
      <ul className={listStyle}>
        {menuItems.map((item, index) => (
          <NavBarItem
            key={item}
            item={item}
            index={index}
            activeIndex={activeIndex}
            isActive={index === activeIndex}
            onItemClicked={handleItemClick}
            backgroundImage={backgroundImages[index]}
          />
        ))}
      </ul>
    </div>
  );
};
