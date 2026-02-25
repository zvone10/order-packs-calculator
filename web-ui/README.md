# Order Packs Calculator - Web UI

A modern React web application for calculating optimal packing orders using Vite, TypeScript, and Zustand state management.

## Features

- **Pack Sizes Management**: Add, update, and remove pack sizes with persistent local storage
- **Order Calculation**: Calculate optimal packing using the backend API
- **Results Display**: View calculation results in a clean tabular format
- **Responsive Design**: Works seamlessly on desktop and mobile devices
- **TypeScript**: Full type safety with TypeScript and TSX
- **State Management**: Simple and efficient state management with Zustand

## Tech Stack

- **React 18** - UI framework
- **Vite** - Fast build tool and dev server
- **TypeScript** - Type-safe JavaScript
- **Zustand** - Lightweight state management
- **CSS** - Modern styling with CSS variables

## Getting Started

### Prerequisites

- Node.js 14+ and npm/yarn

### Installation

```bash
# Install dependencies
npm install
```

### Development

```bash
# Start development server (runs on http://localhost:3000)
npm run dev
```

The dev server includes a proxy to forward API calls to the backend at `http://localhost:8080`.

### Build for Production

```bash
# Build the project
npm run build

# Preview the production build
npm run preview
```

## Project Structure

```
src/
├── components/           # Reusable React components
│   ├── PackSizesInput.tsx       # Pack sizes input and management
│   ├── CalculationSection.tsx   # Calculation input and button
│   └── ResultsTable.tsx         # Results display table
├── store/               # Zustand store
│   └── calculatorStore.ts       # Centralized state management
├── types.ts             # TypeScript type definitions
├── App.tsx              # Main app component
├── App.css              # Application styles
└── main.tsx             # Entry point
```

## API Integration

The app communicates with the backend API at `/api/v1/pack-calculation`.

**Request Format:**
```json
{
  "numberOfItems": 100,
  "boxCapacity": [10, 20, 30]
}
```

**Response Format:**
```json
{
  "totalItems": 100,
  "results": [
    {"capacity": 30, "boxCount": 3},
    {"capacity": 10, "boxCount": 1}
  ]
}
```

## Local Storage

Pack sizes are automatically saved to local storage under the key `packSizes` and restored on app load.

## Best Practices Implemented

- **Component Separation**: Each section has its own component for better maintainability
- **Type Safety**: Full TypeScript coverage for all components and functions
- **State Centralization**: Zustand store for all application state
- **Responsive Design**: Mobile-first approach with media queries
- **Error Handling**: User-friendly error messages and loading states
- **Accessibility**: Semantic HTML and proper form inputs
- **Clean Code**: Simple, readable, and well-organized code