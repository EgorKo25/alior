import {
  Environment,
  Html,
  PresentationControls,
  useGLTF,
} from "@react-three/drei";
import { memo } from "react";
import { useFloating } from "src/features/useFloating";

export const Tablet768: React.FC<{ url: string }> = memo(({ url }) => {
  const tablet = useGLTF("/tablet/scene.gltf");

  useFloating(tablet.scene, {
    range: 0.1,
    speed: 0.005,
    interval: 3000,
  });

  return (
    <>
      <Environment preset="warehouse" />
      <PresentationControls
        global
        polar={[-0.02, 0.02]}
        azimuth={[-0.02, 0.02]}
        config={{ mass: 1, tension: 170, friction: 26 }}
        snap={{ mass: 1, tension: 170, friction: 26 }}
      >
        <primitive
          object={tablet.scene}
          position-y={0.15}
          position-x={-0.25}
          position-z={-4.4}
          rotation={[0, 0.25, 0]}
        >
          <Html
            position={[0.221, 0.185, -0.05]}
            transform
            rotation={[0, -0.25, 0]}
            distanceFactor={1.145}
          >
            <iframe
              src={url}
              className="bg-white w-[768px] h-[1010px] rounded-3xl"
            />
          </Html>
        </primitive>
      </PresentationControls>
    </>
  );
});
