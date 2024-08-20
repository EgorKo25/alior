import { Canvas } from "@react-three/fiber";
import { Suspense } from "react";
import { Tablet1024 } from "./Tablet1024";
import { TabletMobile } from "./TabletMobile";
import { Tv1440 } from "./Tv1440";
import { Tv1920 } from "./Tv1920";
import { Tablet768 } from "./Tablet768";
import { Placeholders } from "./Placeholders";

export const TabletsCanvas = ({ url }: { url: string }) => {
  return (
    <Suspense fallback={<Placeholders />}>
      <Canvas
        id="tablet-mobile"
        camera={{ fov: 45, near: 0.1, far: 2000, position: [0, 0, 0] }}
        className="md:hidden"
      >
        <TabletMobile url={url} />
      </Canvas>
      <Canvas
        id="tablet-768"
        camera={{ fov: 45, near: 0.1, far: 2000, position: [0, 0, 0] }}
        className="hidden md:block lg:hidden"
      >
        <Tablet768 url={url} />
      </Canvas>
      <Canvas
        id="tablet-1024"
        camera={{ fov: 45, near: 0.1, far: 2000, position: [0, 0, 0] }}
        className="hidden lg:block xl:hidden"
      >
        <Tablet1024 url={url} />
      </Canvas>
      <Canvas
        id="tv-1440"
        camera={{ fov: 45, near: 0.1, far: 2000, position: [0, 0, 0] }}
        className="hidden xl:block 2xl:hidden"
      >
        <Tv1440 url={url} />
      </Canvas>
      <Canvas
        id="tv-1920"
        camera={{ fov: 45, near: 0.1, far: 2000, position: [0, 0, 0] }}
        className="hidden 2xl:block"
      >
        <Tv1920 url={url} />
      </Canvas>
    </Suspense>
  );
};
