import { Environment, Html, useGLTF } from "@react-three/drei";
import { memo } from "react";

export const Tablet1024: React.FC<{ url: string }> = memo(({ url }) => {
  const tablet = useGLTF("./tablet/scene.gltf");

  return (
    <>
      <Environment preset="warehouse" />
      <primitive
        object={tablet.scene}
        position-y={0.15}
        position-x={-0.1}
        position-z={-4}
        rotation={[0, 0.25, 0]}
      >
        <Html
          position={[0.208, -0.4, -0.05]}
          transform
          rotation={[0, -0.25, 0]}
          distanceFactor={1.14}
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
