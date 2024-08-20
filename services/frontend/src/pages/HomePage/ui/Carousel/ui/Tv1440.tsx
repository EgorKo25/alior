import { Environment, Html, useGLTF } from "@react-three/drei";
import { memo } from "react";

export const Tv1440: React.FC<{ url: string }> = memo(({ url }) => {
  const tv = useGLTF("./tv/scene.gltf");

  return (
    <>
      <Environment preset="warehouse" />

      <primitive
        object={tv.scene}
        position-y={0}
        position-x={0}
        position-z={-2.25}
        rotation={[0, -0.15, 0]}
      >
        <Html
          position={[-4.55, 0.63, -30]}
          transform
          rotation={[0, 0, 0]}
          distanceFactor={7.8}
        >
          <iframe src={url} className="bg-white w-[1450px] h-[820px]" />
        </Html>
      </primitive>
    </>
  );
});
