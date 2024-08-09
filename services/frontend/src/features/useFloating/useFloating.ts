import { useState, useEffect } from "react";
import { useFrame } from "@react-three/fiber";
import * as THREE from "three";

export const useFloating = (
  ref: THREE.Group,
  options = { range: 0.01, speed: 0.01, interval: 3000 }
) => {
  const { range, speed, interval } = options;
  const [targetRotation, setTargetRotation] = useState([0, 0]);

  useEffect(() => {
    const generateTargetRotation = () => {
      const newAzimuth = THREE.MathUtils.randFloat(-range, range);
      const newPolar = THREE.MathUtils.randFloat(-range, range);
      setTargetRotation([newAzimuth, newPolar]);
    };

    generateTargetRotation(); // set initial target rotation

    const intervalId = setInterval(() => {
      generateTargetRotation();
    }, interval);

    return () => clearInterval(intervalId);
  }, [range, interval]);

  useFrame(() => {
    if (ref) {
      ref.rotation.y = THREE.MathUtils.lerp(
        ref.rotation.y,
        targetRotation[0],
        speed
      );
      ref.rotation.x = THREE.MathUtils.lerp(
        ref.rotation.x,
        targetRotation[1],
        speed
      );
    }
  });
};
