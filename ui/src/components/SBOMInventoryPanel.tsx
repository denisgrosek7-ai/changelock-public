import { useState } from "react";

import { getSBOMImage, searchSBOMComponents } from "../api";
import type { SBOMComponent, SBOMImageResponse } from "../types";

type Props = {
  tenantID?: string;
};

export function SBOMInventoryPanel({ tenantID }: Props) {
  const [componentName, setComponentName] = useState("");
  const [purl, setPURL] = useState("");
  const [imageDigest, setImageDigest] = useState("");
  const [results, setResults] = useState<SBOMComponent[]>([]);
  const [selectedDigest, setSelectedDigest] = useState("");
  const [selectedImage, setSelectedImage] = useState<SBOMImageResponse | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [hasSearched, setHasSearched] = useState(false);

  const trimmedComponentName = componentName.trim();
  const trimmedPURL = purl.trim();
  const trimmedImageDigest = imageDigest.trim();
  const hasSearchFilters = Boolean(trimmedComponentName || trimmedPURL || trimmedImageDigest);

  const exampleQueries = [
    {
      label: "openssl",
      apply: () => {
        setComponentName("openssl");
        setPURL("");
        setImageDigest("");
        setError(null);
      },
    },
    {
      label: "pkg:maven/...",
      apply: () => {
        setComponentName("");
        setPURL("pkg:maven/org.apache.logging.log4j/log4j-core");
        setImageDigest("");
        setError(null);
      },
    },
    {
      label: "sha256 digest",
      apply: () => {
        setComponentName("");
        setPURL("");
        setImageDigest("sha256:");
        setError(null);
      },
    },
  ];

  function normalizeSearchError(searchError: unknown) {
    if (!(searchError instanceof Error)) {
      return "Unable to search SBOM inventory.";
    }
    if (searchError.message.includes("at least one component search filter is required")) {
      return "Start by searching a package name, PURL, or image digest.";
    }
    return searchError.message;
  }

  async function loadImage(digest: string) {
    setSelectedDigest(digest);
    setSelectedImage(await getSBOMImage(digest, 200, tenantID));
  }

  async function handleSearch() {
    setHasSearched(true);
    if (!hasSearchFilters) {
      setError("Start by searching a package name, PURL, or image digest.");
      setResults([]);
      setSelectedDigest("");
      setSelectedImage(null);
      return;
    }
    setLoading(true);
    setError(null);
    try {
      const response = await searchSBOMComponents({
        component_name: trimmedComponentName || undefined,
        purl: trimmedPURL || undefined,
        image_digest: trimmedImageDigest || undefined,
        tenant_id: tenantID || undefined,
        limit: "100",
      });
      setResults(response.components);
      const nextDigest = response.components[0]?.image_digest;
      if (nextDigest) {
        await loadImage(nextDigest);
      } else {
        setSelectedDigest("");
        setSelectedImage(null);
      }
    } catch (searchError) {
      setError(normalizeSearchError(searchError));
      setResults([]);
      setSelectedDigest("");
      setSelectedImage(null);
    } finally {
      setLoading(false);
    }
  }

  return (
    <>
      <section className="panel filters-panel inventory-filters">
        <div className="inventory-header">
          <div>
            <span className="summary-label">SBOM investigation</span>
            <h2>Trace components to digests, documents, and evidence</h2>
            <p>
              Start with a package name, PURL, or digest. This view is optimized for investigation, not raw form entry.
            </p>
          </div>
          <div className="chip-row inventory-example-row">
            {exampleQueries.map((query) => (
              <button key={query.label} className="button button--ghost inventory-example" onClick={query.apply}>
                Try {query.label}
              </button>
            ))}
          </div>
        </div>

        <div className="filters-grid inventory-search-grid">
          <label>
            <span>Component</span>
            <input value={componentName} onChange={(event) => setComponentName(event.target.value)} placeholder="openssl" />
          </label>
          <label>
            <span>PURL</span>
            <input value={purl} onChange={(event) => setPURL(event.target.value)} placeholder="pkg:maven/..." />
          </label>
          <label>
            <span>Image Digest</span>
            <input value={imageDigest} onChange={(event) => setImageDigest(event.target.value)} placeholder="sha256:..." />
          </label>
        </div>
        <p className="inventory-helper">
          Search examples: <code>openssl</code>, <code>pkg:maven/org.apache.logging.log4j/log4j-core</code>, or a full
          image digest.
        </p>
        <div className="filters-actions">
          <button className="button" onClick={() => {
            setComponentName("");
            setPURL("");
            setImageDigest("");
            setResults([]);
            setSelectedDigest("");
            setSelectedImage(null);
            setError(null);
            setHasSearched(false);
          }}>
            Reset
          </button>
          <button className="button button--primary" onClick={() => void handleSearch()} disabled={loading || !hasSearchFilters}>
            {loading ? "Searching…" : "Search Inventory"}
          </button>
        </div>
      </section>

      <section className="content-grid">
        <section className="panel table-panel">
          <div className="table-toolbar">
            <span className="summary-label">SBOM Components</span>
            <strong>{results.length}</strong>
          </div>
          {error ? <div className="panel-empty panel-error">{error}</div> : null}
          {!error && !hasSearched ? (
            <div className="panel-empty inventory-empty-state">
              Start by searching a package name, PURL, or image digest. Example queries are available above to speed up triage.
            </div>
          ) : null}
          {!error && hasSearched && results.length === 0 ? (
            <div className="panel-empty inventory-empty-state">
              No components matched the current query. Try a broader package name or switch to a full digest search.
            </div>
          ) : null}
          {results.length > 0 ? (
            <div className="table-scroll">
              <table className="events-table">
                <thead>
                  <tr>
                    <th>Component</th>
                    <th>Version</th>
                    <th>Type</th>
                    <th>Digest</th>
                  </tr>
                </thead>
                <tbody>
                  {results.map((component) => (
                    <tr
                      key={`${component.id}-${component.image_digest}`}
                      className={selectedDigest === component.image_digest ? "is-selected" : undefined}
                      onClick={() => void loadImage(component.image_digest)}
                    >
                      <td>
                        <div className="event-meta-primary">{component.component_name}</div>
                        {component.purl ? <code className="truncate">{component.purl}</code> : null}
                      </td>
                      <td>{component.component_version || "-"}</td>
                      <td>{component.component_type || "-"}</td>
                      <td>
                        <code className="truncate">{component.image_digest}</code>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          ) : null}
        </section>

        <aside className="panel details-panel">
          {selectedImage ? (
            <>
              <header className="details-header">
                <div>
                  <h2>SBOM Image</h2>
                  <p>{selectedImage.document.image_ref || selectedImage.document.image_digest}</p>
                </div>
                <span className="chip">{selectedImage.document.sbom_format}</span>
              </header>

              <section className="details-section">
                <h3>Document</h3>
                <dl className="details-grid">
                  <dt>Digest</dt>
                  <dd><code>{selectedImage.document.image_digest}</code></dd>
                  <dt>Source Ref</dt>
                  <dd>{selectedImage.document.source_ref || "-"}</dd>
                  <dt>SBOM Hash</dt>
                  <dd><code>{selectedImage.document.sbom_hash || "-"}</code></dd>
                  <dt>Component Count</dt>
                  <dd>{selectedImage.component_count}</dd>
                </dl>
              </section>

              <section className="details-section">
                <h3>Components</h3>
                <ul className="analytics-list inventory-component-list">
                  {selectedImage.components.map((component) => (
                    <li key={component.id}>
                      <div>
                        <strong>{component.component_name}</strong>
                        <div className="event-meta-primary">{component.component_version || "-"}</div>
                      </div>
                      <div className="inventory-component-meta">
                        {component.component_type ? <span className="chip chip--tight">{component.component_type}</span> : null}
                        {component.license ? <span className="chip chip--tight">{component.license}</span> : null}
                      </div>
                    </li>
                  ))}
                </ul>
              </section>
            </>
          ) : (
            <div className="details-empty">
              {hasSearched
                ? "Select a component result to inspect the stored SBOM document, digest, and component inventory."
                : "Run a component, PURL, or digest search first, then inspect the matching SBOM document here."}
            </div>
          )}
        </aside>
      </section>
    </>
  );
}
