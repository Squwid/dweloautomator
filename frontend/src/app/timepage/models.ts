interface Day {
  value: string;
  name: string;
}

interface Keyframe {
  id: string;
  type: string; // either heat cool or off
  day: string; // day of the week
  temperature: number;
}

// what the current status of the temperatures is
interface Current {
  type: string;
  temperature: number;
}
