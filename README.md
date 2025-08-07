# Order Packing API

A Go REST API that solves the optimal order packing problem. Given a customer order quantity, it calculates the most efficient combination of pack sizes to fulfill the order while minimizing total items shipped and number of packs used.

## Problem

Customers can order any number of items, but we only ship complete packs. The challenge is to:
1. Only send whole packs (no breaking open)
2. Minimize total items shipped 
3. Minimize number of packs used

## Solution

The API uses a DFS algorithm to find optimal pack combinations. Default pack sizes: 250, 500, 1000, 2000, 5000 items.

### Examples
- Order 1 item → Ship 1×250 pack
- Order 251 items → Ship 1×500 pack (not 2×250)
- Order 1200 items → Ship 1×1000 + 1×200 packs

## Algorithm Approach

The core algorithm handles several edge cases:

### 1. Exact Matches
If order quantity equals any pack size, return that single pack immediately.

### 2. Orders Smaller Than Smallest Pack
For orders like 100 items when smallest pack is 250, ship the smallest pack (250).

### 3. Orders Larger Than Largest Pack
Use largest packs first, then optimize the remainder:
- Order 12500 → 2×5000 + 1×2000 + 1×500 = 12500

### 4. Complex Combinations
For orders between pack sizes, DFS explores all combinations:
- Order 750 → 1×500 + 1×250 (not 3×250 or 1×1000)
- Order 1200 → 1×1000 + 1×200 (optimal combination)

### 5. Edge Cases Handled
- **Zero/negative orders**: Rejected with validation
- **Large numbers**: Efficiently handles orders up to millions
- **Single pack scenarios**: Optimized path for exact matches
- **Remainder optimization**: When using largest packs, remainder is recalculated optimally

## API

```bash
# Calculate packs for order
GET /api/v1/packs/calculate?orderItemQuantity=1200

# Get all pack sizes
GET /api/v1/packs/sizes

# Add/remove pack sizes
POST /api/v1/packs/sizes
DELETE /api/v1/packs/sizes
```

## Tech Stack

- **Backend**: Go 1.24 + Gin
- **Storage**: Redis (sorted sets)
- **UI**: HTML/CSS/JS
- **Docs**: Swagger
- **Deploy**: Docker

## Architecture

```
Web UI → Gin API → Pack Service → Redis
```

The pack service implements the core algorithm using DFS to find optimal combinations. Redis stores pack sizes in a sorted set for efficient retrieval.

## Testing

Comprehensive test coverage including:
- Exact pack size matches
- Orders smaller than smallest pack
- Orders larger than largest pack  
- Complex combinations requiring multiple pack sizes
- Edge cases and boundary conditions
- Large number scenarios (15,000+ items)

## Quick Start

```bash
docker-compose up -d
```

Access at https://oreder-packing-production.up.railway.app/
