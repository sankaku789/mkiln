addEventListener("DOMContentLoaded", () => {
  const toc = document.getElementById("TOC");
  if (!toc) return;

  const list = toc.querySelector(":scope > ul");
  if (list) {
    list.id = "toc-list";
    const toggle = document.createElement("button");
    toggle.className = "toc-toggle";
    toggle.type = "button";
    toggle.setAttribute("aria-controls", list.id);
    toggle.setAttribute("aria-expanded", "false");
    toggle.innerHTML = '<span aria-hidden="true">☰</span><span>メニュー</span>';
    toggle.addEventListener("click", () => {
      const open = toc.classList.toggle("is-open");
      toggle.setAttribute("aria-expanded", String(open));
    });
    toc.prepend(toggle);

    const mobileHeader = document.createElement("div");
    mobileHeader.className = "mobile-doc-header";
    mobileHeader.append(toc);
    document.body.prepend(mobileHeader);
  }

  const links = [...toc.querySelectorAll('a[href^="#"]')];
  const items = links
    .map((link) => ({ link, heading: document.getElementById(decodeURIComponent(link.hash.slice(1))) }))
    .filter((item) => item.heading);

  if (!items.length) return;

  let current;
  let scheduled = false;

  const update = () => {
    const marker = Math.min(160, window.innerHeight * 0.25);
    const active = items.reduce((selected, item) =>
      item.heading.getBoundingClientRect().top <= marker ? item : selected, items[0]);

    if (active !== current) {
      current?.link.classList.remove("is-current");
      current?.link.removeAttribute("aria-current");
      active.link.classList.add("is-current");
      active.link.setAttribute("aria-current", "location");
      current = active;
    }
    scheduled = false;
  };

  const schedule = () => {
    if (!scheduled) {
      scheduled = true;
      requestAnimationFrame(update);
    }
  };

  addEventListener("scroll", schedule, { passive: true });
  addEventListener("resize", schedule);
  links.forEach((link) => link.addEventListener("click", () => {
    toc.classList.remove("is-open");
    toc.querySelector(".toc-toggle")?.setAttribute("aria-expanded", "false");
  }));
  update();
});
