CREATE TABLE proposals (
    id TEXT PRIMARY KEY  -- hex string (e.g., "0xabc123...")
);

CREATE TABLE proposal_meters (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    proposal_id TEXT NOT NULL,
    m3ter_no INTEGER NOT NULL,
    account TEXT NOT NULL,
    nonce TEXT NOT NULL, -- store big.Int as string
    FOREIGN KEY (proposal_id) REFERENCES proposals(id)
);
