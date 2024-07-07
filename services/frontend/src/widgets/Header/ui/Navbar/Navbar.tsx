import { useState, useEffect } from "react";

import styles from "./Navbar.module.scss";
import mobileStyles from "./MobileNavbar.module.scss";
import { NavbarItem } from "./ui/NavbarItem";

const menuItems = [
  { name: "О нас", sectionId: "features" },
  { name: "Услуги", sectionId: "services" },
  { name: "Кейсы", sectionId: "cases" },
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
    const section = document.getElementById(menuItems[index].sectionId);
    section?.scrollIntoView({ behavior: "smooth" });
  };

  const containerStyle = isMobile
    ? mobileStyles.container
    : styles.container + " bg-accent";
  const listStyle = isMobile ? mobileStyles.list : styles.list;
  return (
    <div className={containerStyle}>
      <ul className={listStyle}>
        {menuItems.map((item, index) => (
          <NavbarItem
            key={item.sectionId}
            item={item.name}
            index={index}
            activeIndex={activeIndex}
            isActive={index === activeIndex}
            onItemClicked={handleItemClick}
          />
        ))}
      </ul>
    </div>
  );
};
