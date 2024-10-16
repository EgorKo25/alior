export const CasesIcon = ({ className }: { className: string }) => {
  return (
    <svg
      className={className}
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <g clipPath="url(#clip0_101_1757)">
        <mask
          id="mask0_101_1757"
          style={{ maskType: "alpha" }}
          maskUnits="userSpaceOnUse"
          x="0"
          y="0"
          width="24"
          height="24"
        >
          <rect
            className={` fill-black lg:fill-white`}
            width="24"
            height="24"
            fill="currentColor"
          />
        </mask>
        <g mask="url(#mask0_101_1757)">
          <path
            className={` fill-black lg:fill-white`}
            d="M11.5001 13.5L17.5001 9.5L11.5001 5.5V13.5ZM12.3251 19H17.8001C17.7001 19.4 17.5001 19.7458 17.2001 20.0375C16.9001 20.3292 16.5334 20.5 16.1001 20.55L5.20012 21.875C4.65012 21.9417 4.15428 21.8083 3.71262 21.475C3.27095 21.1417 3.01678 20.7 2.95012 20.15L1.65012 9.2C1.58345 8.65 1.72095 8.15833 2.06262 7.725C2.40428 7.29167 2.85012 7.04167 3.40012 6.975L4.52512 6.85V8.85L3.62512 8.95L4.97512 19.9L12.3251 19ZM8.52512 17C7.97512 17 7.50428 16.8042 7.11262 16.4125C6.72095 16.0208 6.52512 15.55 6.52512 15V4C6.52512 3.45 6.72095 2.97917 7.11262 2.5875C7.50428 2.19583 7.97512 2 8.52512 2H19.5251C20.0751 2 20.5459 2.19583 20.9376 2.5875C21.3293 2.97917 21.5251 3.45 21.5251 4V15C21.5251 15.55 21.3293 16.0208 20.9376 16.4125C20.5459 16.8042 20.0751 17 19.5251 17H8.52512ZM8.52512 15H19.5251V4H8.52512V15Z"
            fill="currentColor"
          />
        </g>
      </g>
      <defs>
        <clipPath id="clip0_101_1757">
          <rect
            className={` fill-black lg:fill-white`}
            width="24"
            height="24"
            fill="white"
          />
        </clipPath>
      </defs>
    </svg>
  );
};
