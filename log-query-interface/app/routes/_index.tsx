import {LoaderFunctionArgs} from '@remix-run/node'
import {useLoaderData} from '@remix-run/react'
// eslint-disable-next-line import/no-unresolved
import searchLogs from '~/utils/search.server'

export async function loader({request}: LoaderFunctionArgs) {
  const url = new URL(request.url)
  const prompt = url.search

  if (!prompt) return null

  const json = await searchLogs(prompt)
  return new Response(json)
}

export default function Index() {
  const logs = useLoaderData()
  return (
    <div className='flex flex-col h-screen mx-auto'>
      <div className='flex flex-col w-full mx-auto max-w-7xl'>
        <h1 className='text-bold'>Search Results</h1>
      </div>

      <div className='mt-4'>
        {logs &&
          logs.map((log, index) => (
            <div key={index} className='p-2 border rounded mb-2'>
              <pre>{JSON.stringify(log, null, 2)}</pre>
            </div>
          ))}
      </div>
    </div>
  )
}
