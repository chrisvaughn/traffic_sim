<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';
import { GetSimulationState, SetSimulationSpeed, GetSimulationSpeed } from '../../wailsjs/go/main/App';

interface Position {
  x: number;
  y: number;
}

interface Vehicle {
  id: string;
  position: Position;
  velocity: number;
  heading: number;
}

interface StopSign {
  id: string;
  position: Position;
  roadId: string;
}

interface Intersection {
  id: string;
  position: Position;
  stopSigns: StopSign[];
}

interface SimulationState {
  vehicles: Vehicle[];
  intersections: Intersection[];
  time: number;
}

const canvas = ref<HTMLCanvasElement | null>(null);
const currentSpeed = ref(1);
const simTime = ref(0);

let animationFrameId: number;
let ctx: CanvasRenderingContext2D | null = null;

const speedOptions = [1, 2, 5, 10];

const setSpeed = async (speed: number) => {
  await SetSimulationSpeed(speed);
  currentSpeed.value = speed;
};

const drawIntersection = (intersection: Intersection) => {
  if (!ctx) return;

  const { x, y } = intersection.position;

  // Draw intersection as a square
  ctx.fillStyle = '#555';
  ctx.fillRect(x - 30, y - 30, 60, 60);

  // Draw roads extending from intersection
  ctx.fillStyle = '#333';

  // Horizontal road (east-west)
  ctx.fillRect(0, y - 20, canvas.value!.width, 40);

  // Vertical road (north-south)
  ctx.fillRect(x - 20, 0, 40, canvas.value!.height);

  // Draw lane dividers
  ctx.strokeStyle = '#FFD700';
  ctx.lineWidth = 2;
  ctx.setLineDash([10, 10]);

  // Horizontal center line
  ctx.beginPath();
  ctx.moveTo(0, y);
  ctx.lineTo(canvas.value!.width, y);
  ctx.stroke();

  // Vertical center line
  ctx.beginPath();
  ctx.moveTo(x, 0);
  ctx.lineTo(x, canvas.value!.height);
  ctx.stroke();

  ctx.setLineDash([]);
};

const drawStopSign = (stopSign: StopSign) => {
  if (!ctx) return;

  const { x, y } = stopSign.position;

  // Draw stop sign as red octagon
  ctx.fillStyle = '#FF0000';
  ctx.beginPath();
  const size = 8;
  for (let i = 0; i < 8; i++) {
    const angle = (Math.PI / 4) * i;
    const px = x + size * Math.cos(angle);
    const py = y + size * Math.sin(angle);
    if (i === 0) {
      ctx.moveTo(px, py);
    } else {
      ctx.lineTo(px, py);
    }
  }
  ctx.closePath();
  ctx.fill();

  // White border
  ctx.strokeStyle = '#FFFFFF';
  ctx.lineWidth = 1;
  ctx.stroke();
};

const drawVehicle = (vehicle: Vehicle) => {
  if (!ctx) return;

  const { x, y } = vehicle.position;
  // Heading uses standard math coordinates: 0 = East, π/2 = North
  // Canvas rotation also uses standard math coordinates, so no conversion needed

  ctx.save();
  ctx.translate(x, y);
  ctx.rotate(vehicle.heading);

  // Draw vehicle as a rectangle
  ctx.fillStyle = '#4CAF50';
  ctx.fillRect(-10, -5, 20, 10);

  // Draw direction indicator (front of vehicle)
  ctx.fillStyle = '#FFEB3B';
  ctx.fillRect(8, -3, 4, 6);

  ctx.restore();
};

const render = async () => {
  if (!ctx || !canvas.value) return;

  try {
    const state: SimulationState = await GetSimulationState();
    simTime.value = state.time;

    // Clear canvas
    ctx.fillStyle = '#1a1a1a';
    ctx.fillRect(0, 0, canvas.value.width, canvas.value.height);

    // Draw intersections
    state.intersections.forEach(intersection => {
      drawIntersection(intersection);
      intersection.stopSigns.forEach(stopSign => {
        drawStopSign(stopSign);
      });
    });

    // Draw vehicles
    state.vehicles.forEach(vehicle => {
      drawVehicle(vehicle);
    });

  } catch (error) {
    console.error('Error getting simulation state:', error);
  }

  animationFrameId = requestAnimationFrame(render);
};

onMounted(async () => {
  if (canvas.value) {
    ctx = canvas.value.getContext('2d');
    if (ctx) {
      render();
    }
  }

  // Get initial speed
  try {
    currentSpeed.value = await GetSimulationSpeed();
  } catch (error) {
    console.error('Error getting initial speed:', error);
  }
});

onUnmounted(() => {
  if (animationFrameId) {
    cancelAnimationFrame(animationFrameId);
  }
});
</script>

<template>
  <div class="simulator-container">
    <div class="controls">
      <div class="speed-controls">
        <span class="label">Speed:</span>
        <button
          v-for="speed in speedOptions"
          :key="speed"
          :class="{ active: currentSpeed === speed }"
          @click="setSpeed(speed)"
        >
          {{ speed }}x
        </button>
      </div>
      <div class="info">
        <span>Sim Time: {{ simTime.toFixed(1) }}s</span>
      </div>
    </div>
    <canvas
      ref="canvas"
      width="800"
      height="600"
      class="simulation-canvas"
    />
  </div>
</template>

<style scoped>
.simulator-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  gap: 20px;
}

.controls {
  display: flex;
  gap: 30px;
  align-items: center;
  background: #2a2a2a;
  padding: 15px 25px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
}

.speed-controls {
  display: flex;
  gap: 10px;
  align-items: center;
}

.label {
  font-weight: 600;
  color: #ccc;
  margin-right: 5px;
}

button {
  padding: 8px 16px;
  background: #444;
  color: #fff;
  border: 2px solid #666;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  transition: all 0.2s;
}

button:hover {
  background: #555;
  border-color: #777;
}

button.active {
  background: #4CAF50;
  border-color: #4CAF50;
  box-shadow: 0 0 10px rgba(76, 175, 80, 0.5);
}

.info {
  color: #aaa;
  font-size: 14px;
  font-family: monospace;
}

.simulation-canvas {
  border: 2px solid #444;
  border-radius: 4px;
  background: #1a1a1a;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.5);
}
</style>
