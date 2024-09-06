export function incrementCacheVersion() {
  const currentVersion = localStorage.getItem('cacheVersion') || '0';
  const newVersion = (parseInt(currentVersion) + 1).toString();
  localStorage.setItem('cacheVersion', newVersion);
}
