"""Generate GitHub-friendly API docs from pydoc output.

This script walks the ``src`` package, renders pydoc text for each module,
and writes Markdown pages under ``docs/pydoc``.
"""

from __future__ import annotations

from dataclasses import dataclass
from importlib import import_module
from pathlib import Path
import pydoc
import sys


ROOT = Path(__file__).resolve().parents[1]
SRC_ROOT = ROOT / "src"
DOCS_ROOT = ROOT / "docs" / "pydoc"


@dataclass(frozen=True)
class ModuleDoc:
    """Represent one generated module doc page."""

    module_name: str
    output_path: Path


def _iter_source_modules() -> list[str]:
    """Return fully qualified module names for Python source files in src.

    Returns:
        list[str]: Sorted module names under the src package.
    """
    modules: list[str] = []
    for file_path in sorted(SRC_ROOT.rglob("*.py")):
        if "__pycache__" in file_path.parts:
            continue
        if file_path.name == "__init__.py":
            continue

        relative = file_path.relative_to(ROOT).with_suffix("")
        module_name = ".".join(relative.parts)
        modules.append(module_name)
    return modules


def _render_module_doc(module_name: str) -> str:
    """Render plain pydoc output for a module.

    Args:
        module_name: Fully-qualified module name.

    Returns:
        str: Plain text pydoc content.
    """
    module = import_module(module_name)
    return pydoc.plain(pydoc.render_doc(module, title="%s"))


def _module_output_path(module_name: str) -> Path:
    """Map module name to output markdown file path.

    Args:
        module_name: Fully-qualified module name.

    Returns:
        Path: Output markdown path under docs/pydoc.
    """
    return DOCS_ROOT / Path(*module_name.split(".")).with_suffix(".md")


def _write_module_page(module_name: str, output_path: Path, doc_text: str) -> None:
    """Write one module pydoc page as markdown.

    Args:
        module_name: Fully-qualified module name.
        output_path: Markdown file path to write.
        doc_text: Plain text pydoc content.

    Returns:
        None: Writes file content as a side effect.
    """
    output_path.parent.mkdir(parents=True, exist_ok=True)
    output_path.write_text(
        "\n".join(
            [
                f"# {module_name}",
                "",
                "Generated from Python docstrings using `pydoc`.",
                "",
                "```text",
                doc_text.rstrip(),
                "```",
                "",
            ]
        ),
        encoding="utf-8",
    )


def _write_index(pages: list[ModuleDoc]) -> None:
    """Write docs index page linking all generated module docs.

    Args:
        pages: Generated module pages.

    Returns:
        None: Writes index content as a side effect.
    """
    lines = [
        "# PyDoc API Pages",
        "",
        "These pages are generated from project docstrings.",
        "",
        "## Regenerate",
        "",
        "```bash",
        "poetry run python scripts/generate_pydoc_pages.py",
        "```",
        "",
        "## Modules",
        "",
    ]
    for page in pages:
        rel = page.output_path.as_posix().split("/docs/pydoc/", maxsplit=1)[-1]
        lines.append(f"- [{page.module_name}]({rel})")

    index_path = DOCS_ROOT / "README.md"
    index_path.parent.mkdir(parents=True, exist_ok=True)
    index_path.write_text("\n".join(lines) + "\n", encoding="utf-8")


def main() -> int:
    """Generate all pydoc pages.

    Returns:
        int: Process exit code.
    """
    if str(ROOT) not in sys.path:
        sys.path.insert(0, str(ROOT))

    modules = _iter_source_modules()
    pages: list[ModuleDoc] = []
    for module_name in modules:
        output_path = _module_output_path(module_name)
        doc_text = _render_module_doc(module_name)
        _write_module_page(module_name, output_path, doc_text)
        pages.append(ModuleDoc(module_name=module_name, output_path=output_path))

    _write_index(pages)
    print(f"Generated {len(pages)} module pages in {DOCS_ROOT.relative_to(ROOT)}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
