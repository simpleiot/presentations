outline:

- IoT space is exploding right now
- companies/projects popping up everywhere
- reasons for is the availability/lowering cost of key technologies:
  - Connected devices
    - ESP8266/ESP32/Cat-M modems, etc
  - Compute resources at the edge
    - Embedded Linux and SOC devices that run Linux
  - Cloud
  - Connectivity
    - CAT-M/NB-IOT/LORAWAN/BLE Mesh, etc
- many problems to be solved
  - we still can't effectively and efficiently heat and cool most buildings
  - many machines can benefit from some type of monitoring (preventive maint,
    etc)
  - avoid trips to monitor equipment
  - elminate waste and damage
- interested in applications where the system pays for itself
- Video demo of using NATS.io to push data to IoT devices
- useful IoT systems are:
  - realtime
  - distributed
  - malleable
- How Simple IoT is different
  - Single binary with embedded assets
  - Embedded database (may expand to external for large deployments)
  - deploy to cloud or edge devices
  - solving hard problems by simplifying rather than adding layers of complexity
- Data structures
  - Using a common data structure to represent sensor data
    - mechanism for processing and storing the data can stay the same
  - Synchronize config between instances
    - copy the entire config any time it changed
    - inneficient and difficult to sync if multiple things are changing
    - change samples to points, and represent state and config as point. Any
      point can change at any time, and only points are transfered as they
      change.
    - could we represent users, groups and rules as nodes
    - represent the entire system as a tree of nodes
    - users have access to parent node and parent's children
    - rules operate on parent node and children
