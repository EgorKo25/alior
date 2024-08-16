import { Canvas } from "@react-three/fiber";
import { Suspense } from "react";
import { Tablet1024 } from "./Tablet1024";
import { TabletMobile } from "./TabletMobile";
import { Tablet1440 } from "./Tablet1440";
import { Tablet1920 } from "./Tablet1920";
import { Tablet768 } from "./Tablet768";

export const TabletCanvas = ({ url }: { url: string }) => {
  return (
    <Suspense
      fallback={
        <div className=" w-full h-full grid place-items-center">
          <div>Loading...</div>
        </div>
      }
    >
      <Canvas
        camera={{ fov: 45, near: 0.1, far: 2000, position: [0, 0, 0] }}
        className="md:hidden"
      >
        <TabletMobile url={url} />
      </Canvas>
      <Canvas
        camera={{ fov: 45, near: 0.1, far: 2000, position: [0, 0, 0] }}
        className="hidden md:block lg:hidden"
      >
        <Tablet768 url={url} />
      </Canvas>
      <Canvas
        camera={{ fov: 45, near: 0.1, far: 2000, position: [0, 0, 0] }}
        className="hidden lg:block xl:hidden"
      >
        <Tablet1024 url={url} />
      </Canvas>
      <Canvas
        camera={{ fov: 45, near: 0.1, far: 2000, position: [0, 0, 0] }}
        className="hidden xl:block 2xl:hidden"
      >
        <Tablet1440 url={url} />
      </Canvas>
      <Canvas
        camera={{ fov: 45, near: 0.1, far: 2000, position: [0, 0, 0] }}
        className="hidden 2xl:block"
      >
        <Tablet1920 url={url} />
      </Canvas>
    </Suspense>
  );
};
