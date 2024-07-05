export const CircleIcon = () => (
  <div
    className={` circle inline-block w-2.5 h-2.5 rounded-full opacity-100 group-hover:opacity-0 ml-4 
     bg-white 
      group-[.btn-accent-white]:bg-white group-[.btn-accent-white]:hover:bg-accent
      group-[.btn-white-accent]:bg-accent group-[.btn-white-accent]:hover:bg-white
      group-[.btn-orange-white]:hover:bg-orange-900
      group-[.btn-blue-white]:hover:bg-blue-900
      group-[.btn-black-white]:hover:bg-black
      group-[.btn-white-black]:bg-black group-[.btn-white-black]:hover:bg-white`}
    style={{ transition: "opacity 0.3s" }}
  ></div>
);
