import { Environment, Html, useGLTF } from "@react-three/drei";
import { memo } from "react";

export const Tv1920: React.FC<{ url: string }> = memo(({ url }) => {
  const tv = useGLTF("./tv/scene.gltf");

  return (
    <>
      <Environment preset="warehouse" />

      <primitive
        object={tv.scene}
        position-y={0.05}
        position-x={0}
        position-z={-2}
        rotation={[0, -0.15, 0]}
      >
        <Html
          position={[-7.6, 1.77, -50]}
          transform
          rotation={[0, 0, 0]}
          distanceFactor={14.15}
        >
          <iframe src={url} className="bg-white w-[1450px] h-[820px]" />
        </Html>
      </primitive>
    </>
  );
});
