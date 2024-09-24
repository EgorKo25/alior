import { Link } from "react-router-dom";
import styles from "./NavbarItem.module.scss";
import clsx from "clsx";
import { ServicesIcon } from "./ui/ServicesIcon";
import { AboutIcon } from "./ui/AboutIcon";
import { CasesIcon } from "./ui/CasesIcon";

interface NavigationItemProps {
  item: string;
  index: number;
  onItemClicked: (index: number) => void;
}

const menuItems = [
  { name: "О нас", sectionId: "features" },
  { name: "Услуги", sectionId: "services" },
  { name: "Кейсы", sectionId: "cases" },
];

export const NavbarItem: React.FC<NavigationItemProps> = ({
  item,
  index,
  onItemClicked,
}) => {
  const icons = [
    <AboutIcon className={clsx(styles.icon)} />,
    <ServicesIcon className={clsx(styles.icon)} />,
    <CasesIcon className={clsx(styles.icon)} />,
  ];
  return (
    <li
      className={clsx(styles.list__item) + " 2xl:py-4 2xl:px-5"}
      onClick={() => onItemClicked(index)}
    >
      <div className={clsx(styles.decoration_container)}>
        {icons[index]}
        <div className={clsx(styles.decoration)}></div>
      </div>
      <Link
        to={{ pathname: "/", hash: `#${menuItems[index].sectionId}` }}
        className={
          clsx(styles.link) +
          " sm:text-black md:text-black lg:text-white 2xl:text-xl"
        }
      >
        {item}
      </Link>
    </li>
  );
};
