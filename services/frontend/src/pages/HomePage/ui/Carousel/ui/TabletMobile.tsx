import { Environment, Html, useGLTF } from "@react-three/drei";
import { memo } from "react";

export const TabletMobile: React.FC<{ url: string }> = memo(({ url }) => {
  const tablet = useGLTF("./tablet/scene.gltf");

  return (
    <>
      <Environment preset="warehouse" />
      <primitive
        object={tablet.scene}
        position-y={0.15}
        position-x={-0.2}
        position-z={-4.2}
        rotation={[0, 0.25, 0]}
      >
        <Html
          position={[0.225, 0, -0.105]}
          transform
          rotation={[0, -0.25, 0]}
          distanceFactor={1.15}
        >
          <iframe
            src={url}
            className="bg-white w-[768px] h-[1010px] rounded-3xl"
          />
        </Html>
      </primitive>
    </>
  );
});
