import { useEffect, useRef, useState } from "react";
import "./style.scss";

export const IframeScreen: React.FC<{ url: string }> = ({ url }) => {
  const [iframeScale, setIframeScale] = useState(0);
  const [windowWidth, setWindowWidth] = useState(window.innerWidth);
  const iframeWrapper = useRef<HTMLDivElement>(null);
  const iframe = useRef<HTMLIFrameElement>(null);

  useEffect(() => {
    window.addEventListener("resize", () => {
      setWindowWidth(window.innerWidth);
    });
  }, []);

  useEffect(() => {
    if (iframeWrapper.current && iframe.current) {
      setIframeScale(
        iframeWrapper.current.clientWidth / iframe.current.clientWidth
      );
    }
  }, [windowWidth]);
  return (
    <div className="screen-wrapper w-full h-full grid place-items-center max-w-[1020px]">
      <div className="screen-parent relative w-full lg:w-[85%] xl:w-full aspect-[9/16] md:aspect-[4/5] xl:aspect-video">
        <div className="screen-slice-1 z-[1]"></div>
        <div className="screen-slice-2 z-[2]"></div>
        <div className="screen-slice-3 z-[3]"></div>
        <div className="screen-slice-4 z-[4]"></div>
        <div className="screen-slice-5 z-[5]"></div>
        <div className="screen-slice-6 z-[6]"></div>
        <div className="screen-slice-7 z-[7]"></div>
        <div className="screen-slice-8 z-[8]"></div>
        <div className="screen-slice-9 z-[9]"></div>
        <div className="screen-slice-10 z-[10]"></div>
        <div className="screen-slice-11 z-[11]"></div>
        <div className="screen-slice-12 z-[13]"></div>
        <div className="screen-slice-13 z-[12]"></div>
        <div className="screen-slice-14 z-[14]"></div>
        <div className="screen-slice-15 z-[15]"></div>
        <div className="screen-slice-16 z-[16]"></div>
        <div className="screen-slice-17 z-[17]"></div>
        <div className="screen-slice-18 z-[18]"></div>
        <div className="screen-slice-19 z-[19]"></div>
        <div className="screen-slice-20 z-[20]"></div>
        <div className="screen-slice-21 z-[21]"></div>
        <div className="screen-slice-22 z-[22]"></div>
        <div className="screen-slice-23 z-[23]"></div>
        <div className="screen-slice-24 z-[24]"></div>
        <div className="screen-slice-25 z-[25]"></div>
        <div className="screen-slice-26 z-[26]">
          <div
            className="iframe-wrapper w-full h-full overflow-hidden relative rounded-20"
            ref={iframeWrapper}
          >
            <iframe
              src={url}
              className="iframe-content lg:rounded-40 border-solid md:border-8 border-transparent w-[450px] md:w-[1024px] xl:w-[1440px] aspect-[9/16] md:aspect-[4/5] xl:aspect-video absolute left-0 top-0"
              style={{
                transform: `scale(${iframeScale})`,
                transformOrigin: "top left",
              }}
              ref={iframe}
            ></iframe>
          </div>
        </div>
      </div>
    </div>
  );
};
