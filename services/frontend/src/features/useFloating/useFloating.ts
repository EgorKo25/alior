import { useState, useEffect, useRef } from "react";
import { useFrame } from "@react-three/fiber";
import * as THREE from "three";

export const useFloating = (
  ref: THREE.Group,
  options = { range: 0.01, speed: 0.01 }
) => {
  const { range, speed } = options;

  const initialRotation = useRef<[number, number]>([0, 0]);

  const [targetRotation, setTargetRotation] = useState<[number, number]>([
    0, 0,
  ]);

  const [shouldAnimate, setShouldAnimate] = useState(true);

  useEffect(() => {
    if (ref) {
      initialRotation.current = [ref.rotation.x, ref.rotation.y];
    }
  }, [ref]);

  useFrame(() => {
    if (!ref) return;

    if (shouldAnimate) {
      const generateTargetRotation = () => {
        const newAzimuth = THREE.MathUtils.randFloat(-range, range);
        const newPolar = THREE.MathUtils.randFloat(-range, range);
        setTargetRotation([newAzimuth, newPolar]);
      };

      generateTargetRotation();
      setShouldAnimate(false);
    }

    ref.rotation.y = THREE.MathUtils.lerp(
      ref.rotation.y,
      initialRotation.current[1] + targetRotation[0],
      speed
    );
    ref.rotation.x = THREE.MathUtils.lerp(
      ref.rotation.x,
      initialRotation.current[0] + targetRotation[1],
      speed
    );

    const reachedTarget =
      Math.abs(
        ref.rotation.y - (initialRotation.current[1] + targetRotation[0])
      ) < 0.005 &&
      Math.abs(
        ref.rotation.x - (initialRotation.current[0] + targetRotation[1])
      ) < 0.005;

    if (reachedTarget) {
      setShouldAnimate(true);
    }
  });
};
