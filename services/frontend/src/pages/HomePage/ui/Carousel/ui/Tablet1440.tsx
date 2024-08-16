import {
  Environment,
  Html,
  PresentationControls,
  useGLTF,
} from "@react-three/drei";
import { memo } from "react";
import { useFloating } from "src/features/useFloating";

export const Tablet1440: React.FC<{ url: string }> = memo(({ url }) => {
  const tablet = useGLTF("/tablet/scene.gltf");

  useFloating(tablet.scene, {
    range: 0.02,
    speed: 0.005,
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
          position-y={-0.1}
          position-x={0.1}
          position-z={-3.8}
          rotation={[-0.4, -0.15, Math.PI * 0.45]}
        >
          <Html
            position={[0.62, 0.11, -0.55]}
            transform
            rotation={[0, -0.25, Math.PI * -0.5]}
            distanceFactor={1.042}
          >
            <iframe
              src={url}
              className="bg-white w-[1280px] h-[980px] rounded-40"
            />
          </Html>
        </primitive>
      </PresentationControls>
    </>
  );
});
