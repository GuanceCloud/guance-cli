import { ChevronRightIcon } from '@heroicons/react/20/solid'

export default function Example() {
    return (
        <div className="bg-white">
            <div className="relative isolate overflow-hidden bg-gradient-to-b from-indigo-100/20">
                <div className="mx-auto max-w-7xl pb-24 pt-10 sm:pb-32 lg:grid lg:grid-cols-2 lg:gap-x-8 lg:px-8 lg:py-40">
                    <div className="px-6 lg:px-0 lg:pt-4">
                        <div className="mx-auto max-w-2xl">
                            <div className="max-w-lg">
                                <img
                                    className="h-11"
                                    src="/icon.png"
                                    alt="Guance Cloud"
                                />
                                <div className="mt-24 sm:mt-32 lg:mt-16">
                                    <a href="/docs/how-to-guides/grafana" className="inline-flex space-x-6">
                    <span className="rounded-full bg-indigo-600/10 px-3 py-1 text-sm font-semibold leading-6 text-indigo-600 ring-1 ring-inset ring-indigo-600/10">
                      What's new
                    </span>
                                        <span className="inline-flex items-center space-x-2 text-sm font-medium leading-6 text-gray-600">
                      <span>
                            🎉 Grafana importer is released.
                      </span>
                      <ChevronRightIcon className="h-5 w-5 text-gray-400" aria-hidden="true" />
                    </span>
                                    </a>
                                </div>
                                <h1 className="mt-10 text-4xl font-bold tracking-tight text-gray-900 sm:text-6xl">
                                    Your toolkit for Guance Cloud control-plane
                                </h1>
                                <p className="mt-6 text-lg leading-8 text-gray-600">
                                    Developer-first toolkit to manage / export / import cloud resources on Guance.
                                </p>
                                <div className="mt-10 flex items-center gap-x-6">
                                    <a
                                        href="/features"
                                        className="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                                    >
                                        View all features
                                    </a>
                                    <a href="/docs/tutorials/quickstart" className="text-sm font-semibold leading-6 text-gray-900">
                                        Read documentation <span aria-hidden="true">→</span>
                                    </a>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="mt-20 sm:mt-24 md:mx-auto md:max-w-2xl lg:mx-0 lg:mt-0 lg:w-screen">
                        <div className="mockup-code min-h-full">
                            <pre data-prefix="$">guance login</pre>
                            <pre data-prefix=">" className="text-success">Login as bot@guance.com!</pre>
                            <pre></pre>
                            <pre data-prefix="#" className="text-neutral-500">Import from external system</pre>
                            <pre data-prefix="$">guance iac import grafana --search --search-tags o11y</pre>
                            <pre data-prefix=">" className="text-warning">Found 42 dashboards.</pre>
                            <pre data-prefix=">" className="text-success">Generate Guance code files.</pre>
                            <pre></pre>
                            <pre data-prefix="#" className="text-neutral-500">Manage cloud resources</pre>
                            <pre data-prefix="$">guance api create-resource pipeline ...</pre>
                            <pre data-prefix=">" className="text-success">Done!</pre>
                        </div>
                    </div>
                </div>
                <div className="absolute inset-x-0 bottom-0 -z-10 h-24 bg-gradient-to-t from-white sm:h-32" />
            </div>
        </div>
    )
}
