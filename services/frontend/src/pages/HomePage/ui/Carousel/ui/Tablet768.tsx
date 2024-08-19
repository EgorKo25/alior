import { Environment, Html, useGLTF } from "@react-three/drei";
import { memo } from "react";

export const Tablet768: React.FC<{ url: string }> = memo(({ url }) => {
  const tablet = useGLTF("./tablet/scene.gltf");

  return (
    <>
      <Environment preset="warehouse" />
      <primitive
        object={tablet.scene}
        position-y={0.05}
        position-x={-0.25}
        position-z={-4.4}
        rotation={[0, 0.25, 0]}
      >
        <Html
          position={[6.23, 0.45, -25]}
          transform
          rotation={[0, -0.25, 0]}
          distanceFactor={7.65}
        >
          <iframe
            src={url}
            className="bg-white w-[768px] h-[1010px] rounded-[32px]"
          />
        </Html>
      </primitive>
    </>
  );
});
