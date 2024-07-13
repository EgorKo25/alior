import { Link } from "react-router-dom";
import styles from "./NavbarItem.module.scss";
import clsx from "clsx";
import { ServicesIcon } from "../ServicesIcon";
import { AboutIcon } from "../AboutIcon";
import { CasesIcon } from "../CasesIcon";

interface NavigationItemProps {
  item: string;
  index: number;
  isActive: boolean;
  onItemClicked: (index: number) => void;
  activeIndex: number;
}

const menuItems = [
  { name: "О нас", sectionId: "features" },
  { name: "Услуги", sectionId: "services" },
  { name: "Кейсы", sectionId: "cases" },
];

export const NavbarItem: React.FC<NavigationItemProps> = ({
  item,
  index,
  activeIndex,
  isActive,
  onItemClicked,
}) => {
  const icons = [
    <AboutIcon
      className={clsx(styles.icon, {
        [styles.icon_active]: index === activeIndex,
      })}
    />,
    <ServicesIcon
      className={clsx(styles.icon, {
        [styles.icon_active]: index === activeIndex,
      })}
    />,
    <CasesIcon
      className={clsx(styles.icon, {
        [styles.icon_active]: index === activeIndex,
      })}
    />,
  ];
  return (
    <li
      className={clsx(styles.list__item) + " 2xl:py-4 2xl:px-5"}
      onClick={() => onItemClicked(index)}
    >
      <div className={clsx(styles.decoration_container)}>
        {icons[index]}
        <div
          className={clsx(styles.decoration, {
            " bg-black sm:bg-black md:bg-black lg:bg-white": !isActive,
            " bg-transparent sm:fill-black md:fill-black": isActive,
            [styles.decoration_active]: index === activeIndex,
          })}
        ></div>
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
