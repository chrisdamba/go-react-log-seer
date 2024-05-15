import * as React from 'react'
import {
  Form,
  Links,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
} from '@remix-run/react'
import type {LinksFunction} from '@remix-run/node'

import appStylesHref from './app.css?url'

export const links: LinksFunction = () => [
  {rel: 'stylesheet', href: appStylesHref},
]

export default function App() {
  const [filters, setFilters] = React.useState({
    fullTextSearch: '',
    level: '',
    message: '',
    resourceId: '',
    timestampStart: '',
    timestampEnd: '',
    traceId: '',
    spanId: '',
    commit: '',
    parentResourceId: '',
  })

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const {name, value} = e.target
    setFilters({...filters, [name]: value})
  }

  return (
    <html lang='en'>
      <head>
        <meta charSet='utf-8' />
        <meta name='viewport' content='width=device-width, initial-scale=1' />
        <Meta />
        <Links />
      </head>
      <body>
        <div id='sidebar'>
          <div>
            <Form id='search-form' role='search' method='get'>
              <input
                id='q'
                aria-label='Search contacts'
                placeholder='Search'
                type='search'
                name='fullTextSearch'
              />
              <div id='search-spinner' aria-hidden hidden={true} />
              <div className='mb-4' />
              <div className='grid grid-cols-1 md:grid-cols-2 gap-4 mb-4'>
                {[
                  'level',
                  'message',
                  'resourceId',
                  'timestampStart',
                  'timestampEnd',
                  'traceId',
                  'spanId',
                  'commit',
                  'parentResourceId',
                ].map((key) => (
                  <input
                    key={key}
                    name={key}
                    placeholder={key.charAt(0).toUpperCase() + key.slice(1)}
                    value={filters[key]}
                    onChange={handleInputChange}
                    className='p-2 border rounded'
                  />
                ))}
              </div>
              <button type='submit'>Search</button>
            </Form>
          </div>
        </div>
        <Outlet />
        <ScrollRestoration />
        <Scripts />
      </body>
    </html>
  )
}
