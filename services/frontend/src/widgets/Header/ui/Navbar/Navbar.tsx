import styles from "./Navbar.module.scss";
import mobileStyles from "./MobileNavbar.module.scss";
import { NavbarItem } from "./ui/NavbarItem";

const menuItems = [
  { name: "О нас", sectionId: "features" },
  { name: "Услуги", sectionId: "services" },
  { name: "Кейсы", sectionId: "cases" },
];

export const Navbar = ({ isMobile }: { isMobile?: boolean }) => {
  const handleItemClick = (index: number) => {
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
            onItemClicked={handleItemClick}
          />
        ))}
      </ul>
    </div>
  );
};
